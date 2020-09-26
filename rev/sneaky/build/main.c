#include <stdio.h>
#include <stdlib.h>
#include <time.h>
#include <unistd.h>
#include <sys/ptrace.h>
#include <sys/types.h>
#include <sys/wait.h>
#include <signal.h>
#include <ncurses.h>

#define FLAG_SCORE 10000

#define DIR_RIGHT 1
#define DIR_LEFT  2
#define DIR_UP    3
#define DIR_DOWN  4

#define APPLE_TIMER 10
#define SPEED_INIT 200000
#define SPEED_MAX  50000
#define WIN_WIDTH  22
#define WIN_HEIGHT 22
#define C_GAMEOVER 1
#define C_SNAKE    2
#define C_APPLE    3
#define C_BG       4

#define ROR(c, x) ((c >> x) | ((c & ((1<<x)-1)) << (16-x)))

typedef struct _Body {
  int px, py;
  struct _Body *next;
} Body;
typedef struct {
  Body *head;
  int timer;
  int apple;
  int ax, ay;
  int speed;
  int dir;
  long score;
} Game;

unsigned short list[] = {32003, 46224, 49079, 21155, 60687, 9475, 63426, 34582, 34680, 12198, 590, 56428, 16973};

void anti_debug(void) {
  int pid = fork();
  if (pid == -1) {
    _exit(0);
  } else if (pid != 0) {
    int status;
    signal(SIGINT, SIG_IGN);
    if (ptrace(PTRACE_ATTACH, pid, NULL, NULL) < 0)
      kill(pid, SIGSEGV);
    waitpid(pid, &status, 0);
    ptrace(PTRACE_CONT, pid, NULL, NULL);
    while(waitpid(pid, &status, 0)) {
      if (WIFSTOPPED(status)) {
        ptrace(PTRACE_CONT, pid, NULL, NULL);
      } else {
        _exit(0);
      }
    }
  }
}

__attribute__((constructor))
void setup(void) {
  int wx, wy;
  anti_debug();
  initscr();
  getmaxyx(stdscr, wx, wy);
  if (wx < WIN_WIDTH || wy < WIN_HEIGHT) {
    puts("Terminal size is too small");
    exit(1);
  }
  start_color();
  init_pair(C_GAMEOVER, COLOR_RED, COLOR_WHITE);
  init_pair(C_SNAKE   , COLOR_GREEN, COLOR_BLACK);
  init_pair(C_APPLE   , COLOR_CYAN, COLOR_BLACK);
  init_pair(C_BG      , COLOR_WHITE, COLOR_BLACK);
  keypad(stdscr, TRUE);
  timeout(0);
  noecho();
  curs_set(0);
  srand(time(NULL));
}
__attribute__((destructor))
void cleanup(void) {
  endwin();
}

unsigned short hash(unsigned char *s) {
  unsigned short h = 0x3141;
  h = (h * 0xcafe + s[0]) & 0xffffffff;
  h = (h * 0xdead + s[1]) & 0xffffffff;
  return h;
}
char *gen_flag(Game *game) {
  int i;
  unsigned int block;
  char *flag = malloc(0x40);
  char x = game->score / FLAG_SCORE > 0;
  for(i = 0; i < 13; i++) {
    for(block = 0; block <= 0xffff; block++) {
      if (block == 0x48e8) continue;
      if (hash((unsigned char*)&block) == ROR(list[i], x)) break;
    }
    *(unsigned short*)&flag[i*2] = block;
  }
  return flag;
}

void sneaky_add_body(Game *game, int x, int y) {
  Body *body;
  for(body = game->head; body->next; body = body->next);
  body->next = (Body*)malloc(sizeof(Body));
  body->next->px = x;
  body->next->py = y;
  body->next->next = NULL;
}

void sneaky_update(Game *game) {
  Body *body, *prev;
  int opx, opy, px, py;
  if (game->dir == 0) return;

  for(opx = game->head->px, opy = game->head->py, body = game->head->next;
      body != NULL;
      body = body->next) {
    px = body->px;
    py = body->py;
    body->px = opx;
    body->py = opy;
    opx = px;
    opy = py;
  }
  
  switch(game->dir) {
  case DIR_RIGHT: game->head->px += 1; break;
  case DIR_LEFT : game->head->px -= 1; break;
  case DIR_UP   : game->head->py -= 1; break;
  case DIR_DOWN : game->head->py += 1; break;
  }

  if (!game->apple) {
    if (game->timer <= 0) {
      game->apple = 1;
      game->ax = 1 + rand() % (WIN_WIDTH - 2);
      game->ay = 2 + rand() % (WIN_HEIGHT - 2);
    } else {
      game->timer--;
    }
  }

  if (game->head->px == game->ax && game->head->py == game->ay) {
    sneaky_add_body(game, opx, opy);
    game->apple = 0;
    game->timer = APPLE_TIMER;
    game->score += 10;
    game->speed -= 10000;
    if (game->speed < SPEED_MAX) {
      game->speed = SPEED_MAX;
    }
  }
}

void sneaky_draw(Game *game) {
  Body *body;
  attron(COLOR_PAIR(C_SNAKE));
  for(body = game->head; body; body = body->next) {
    move(body->py, body->px);
    addch('O');
  }

  if (game->apple) {
    attron(COLOR_PAIR(C_APPLE));
    move(game->ay, game->ax);
    addch('$');
  }
}

int sneaky_check(Game *game) {
  Body *body, *target;

  if (game->head->px < 1 || game->head->px > WIN_WIDTH - 2) goto GAMEOVER;
  if (game->head->py < 2 || game->head->py > WIN_HEIGHT - 1) goto GAMEOVER;

  for(body = game->head; body->next; body = body->next) {
    for(target = body->next; target; target = target->next) {
      if (target->px == body->px && target->py == body->py) goto GAMEOVER;
    }
  }

  if (game->score >= FLAG_SCORE) {
    attron(COLOR_PAIR(C_GAMEOVER));
    move(0, 0);
    addstr(gen_flag(game));
    refresh();
    timeout(-1);
    getch();
    return 1;
  }
  return 0;

 GAMEOVER:
  attron(COLOR_PAIR(C_GAMEOVER));
  move(WIN_HEIGHT / 2, WIN_WIDTH / 2 - 8);
  addstr("*** GAMEOVER ***");
  refresh();
  timeout(-1);
  getch();
  return 1;
}

void sneaky(Game *game) {
  short c = 0;
  WINDOW *w = subwin(stdscr, WIN_WIDTH, WIN_HEIGHT, 1, 0);
  wbkgd(w, COLOR_PAIR(C_BG));
  while(c != 27) {
    wclear(w);
    attron(COLOR_PAIR(C_BG));
    move(0, 0);
    printw("SCORE: %d", game->score);
    box(w, 0, 0);
    sneaky_update(game);
    sneaky_draw(game);
    if (sneaky_check(game)) {
      break;
    }
    wrefresh(w);
    usleep(game->speed);

    c = getch();
    switch(c) {
    case KEY_RIGHT: game->dir = DIR_RIGHT; ;break;
    case KEY_LEFT : game->dir = DIR_LEFT;  break;
    case KEY_UP   : game->dir = DIR_UP;    break;
    case KEY_DOWN : game->dir = DIR_DOWN;  break;
    }
    flushinp();
  }
  delwin(w);
}

int main(int argc, char **argv) {
  Game *game = (Game*)calloc(sizeof(Game), 1);
  game->head = (Body*)malloc(sizeof(Body));
  game->head->px = (WIN_WIDTH - 2) / 2;
  game->head->py = (WIN_HEIGHT - 2) / 2 + 1;
  game->head->next = NULL;
  
  game->dir = 0;
  game->speed = SPEED_INIT;
  game->score = 0;
  game->ax = game->ay = 0;
  game->apple = 0;
  game->timer = APPLE_TIMER;
  sneaky(game);
}

import os
import random
import time
from hashlib import sha256

from ariadne import ObjectType, graphql_sync, make_executable_schema, format_error
from flask import Flask, request, jsonify, session, redirect, render_template
import tinydb


type_defs = """
    type User {
        uid: ID!
        name: String!
        profile: String!
        token: String!
    }

    type Query {
        me: User!
        someone(uid: ID!): User
        flag: String!
    }

    type Mutation {
        updateProfile(profile: String!, token: String!): Boolean!
        updateToken(token: String!): Boolean!
    }
"""

query = ObjectType("Query")
mutation = ObjectType("Mutation")


@query.field("me")
def resolve_me(obj, info):
    uid = info.context["uid"]

    result = db.search(db_query.uid == uid)
    if len(result) == 0:
        return 

    user = result[0]
    return {
        "uid": user["uid"],
        "name": user["name"],
        "profile": user["profile"],
        "token": user["token"]
    }


@query.field("someone")
def resolve_someone(obj, info, uid):
    result = db.search(db_query.uid == uid)

    if len(result) == 0:
        return 

    user = result[0]
    return {
        "uid": user["uid"],
        "name": user["name"],
        "profile": user["profile"],
        "token": user["token"]
    }


@query.field("flag")
def resolve_flag(obj, info):
    if info.context.get("uid", None) is None:
        return "Please register your account."

    uid = info.context["uid"]

    result = db.search(db_query.uid == uid)
    if len(result) == 0:
        return "Please register your account."
    
    token = result[0]["token"]
    if token == os.getenv("CTF4B_ADMIN_TOKEN"):
        return os.getenv("CTF4B_FLAG")
    else:
        return "Sorry, your token is not administrator's one. This page is only for administrator(uid: admin)."


@mutation.field("updateProfile")
def resolve_update_profile(obj, info, profile, token):
    if info.context.get("uid", None) is None:
        return False

    uid = info.context["uid"]

    result = db.search(db_query.uid == uid)
    if len(result) == 0:
        return False

    if result[0]["token"] != token:
        return False

    if len(profile) > 1000:
        return False

    try:
        db.update({"profile": profile}, db_query.uid == uid)
        return True
    except:
        return False


@mutation.field("updateToken")
def resolve_update_token(obj, info, token):
    if info.context.get("uid", None) is None:
        return False

    uid = info.context["uid"]

    result = db.search(db_query.uid == uid)
    if len(result) == 0:
        return

    try:
        db.update({"token": token}, db_query.uid == uid)
        return True
    except:
        return False

def init_db():
    global db
    global db_query

    db = tinydb.TinyDB("db.json")
    db.purge()
    db.insert({
        "uid": "admin",
        "password": os.getenv("CTF4B_ADMIN_PASSWORD"),
        "name": "admin",
        "profile": "Hello, I'm admin.",
        "token": os.getenv("CTF4B_ADMIN_TOKEN")
    })
    db_query = tinydb.Query()

app = Flask(__name__)
app.secret_key = os.getenv("CTF4B_SECRET_KEY")
db = None
db_query = None
init_db()

@app.route("/")
def index():
    if "uid" in session:
        return redirect("/profile")
    return render_template("index.html")


@app.route("/register", methods=["POST"])
def register():
    if "uid" in session:
        return redirect("/profile")

    uid = request.form["uid"]
    password = request.form["password"]
    name = request.form["name"]

    if db.contains(db_query.uid == uid):
        return render_template("index.html", error_message="Already registered.")

    if len(uid) == 0 or len(password) == 0 or len(name) == 0:
        return render_template("index.html", error_message="Fill out all fields.")

    if len(uid) > 100 or len(name) > 100:
        return render_template("index.html", error_message="Too long (len <= 100).")

    token = sha256(str(random.random()).encode()).hexdigest()
    db.insert({
        "uid": uid,
        "password": sha256(password.encode()).hexdigest(),
        "name": name,
        "profile": "",
        "token": token
    })

    return render_template("index.html", info_message=f"""Registered successfully. 
    Your token is {token}. Don't lose it!""")


@app.route("/login", methods=["POST"])
def login():
    if "uid" in session:
        return redirect("/profile")

    uid = request.form["uid"]
    password = request.form["password"]

    result = db.search(db_query.uid == uid)
    if len(result) == 0:
        return render_template("index.html", error_message="Login failed.")

    user = result[0]
    if user["password"] != sha256(password.encode()).hexdigest():
        return render_template("index.html", error_message="Login failed.")

    session["uid"] = uid
    return redirect("/profile")


@app.route("/profile")
def profile():
    if "uid" not in session:
        return render_template("index.html", error_message="You are not logged in.")

    return render_template("profile.html")


@app.route("/flag")
def flag():
    return render_template("flag.html")


@app.route("/logout", methods=["POST"])
def logout():
    if "uid" in session:
        session.pop("uid", None)
    return redirect("/")


@app.route("/api", methods=["POST"])
def graphql_server():
    success, result = graphql_sync(
        schema=make_executable_schema(type_defs, query, mutation),
        data=request.get_json(),
        context_value=session,
    )

    status_code = 200 if success else 400
    return jsonify(result), status_code

if __name__ == "__main__":
    init_db()
    app.run(host=os.getenv("CTF4B_HOST"), port=os.getenv("CTF4B_PORT"))

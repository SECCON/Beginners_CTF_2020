class Account:
    def __init__(self, name, password):
        self.name = name
        self.password = password


accounts = None
employees = None

def init():
    global accounts
    global employees

    accounts = [
        Account("Elbert", "impossible_value"),
        Account("George", "impossible_value"),
        Account("Lazarus", "impossible_value"),
        Account("Marc", "impossible_value"),
        Account("Tony", "impossible_value"),
        Account("Ximena", "impossible_value"),
        Account("Yvonne", "impossible_value")
    ]

    employees = [
        "Arthur", "Barbara", "Christine", "David", "Elbert", 
        "Franklin", "George", "Harris", "Ivan", "Jane", 
        "Kevin", "Lazarus", "Marc", "Nathan", "Oliver", 
        "Paul", "Quentin", "Randolph", "Scott", "Tony", 
        "Ulysses", "Vincent", "Wat", "Ximena", "Yvonne", "Zalmon"
    ]


def get_all_accounts():
    return accounts


def get_all_employees():
    return employees


def get_account(name):
    for account in accounts:
        if account.name == name:
            return (True, account)
    
    return (False, None)
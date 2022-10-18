# Budgetly

A simpler way to manage your manual budgeting.

## TODO:

- Integrate DB
    - ORM
        - https://github.com/go-gorm/gorm

Main View
----------

Account
    - Name
    - AccountType
    - Amount

Revenue
    - Amount
    - Account
    - Recurring (bool)
    - DepositDate

Expense
    - Name
    - Type (auto loan, mortgage, cell phone, etc)
    - Category (utilities, entertainment, food, etc)
    - Amount
    - Recurring (bool)
    - DueDate
    - Paid (bool)

- API
    -

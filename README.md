# ama-billing

## Background
This is a Billing System that means to do several things:
1. Loan schedule for a given loan (when am i supposed to pay how much)
2. Outstanding Amount for a given loan
3. Status of whether the customer is Delinquent or not

With above capabilities, this service are meant to at least have these capabilities/goals:
1. GetOutstanding: This returns the current outstanding on a loan, 0 if no outstanding (or closed)
2. IsDelinquent: If there are more than 2 weeks of Non payment of the loan amount
3. MakePayment: Make a payment of certain amount on the loan

## Assumptions
To simplify the process, some assumptions that is applied to the System was:
1. one user ID can only have one active loan at a time

## Run the Service
Use the following command:
```
$ go mod tidy
$ go run .
```

## Capabilities
To fulfill the above requirements, this Billing System have several capabilities:
- Register a Loan
```
curl --location 'localhost:8080/loan' \
--header 'Content-Type: application/json' \
--data '{
    "user_id": 10,
    "principal_value": 100000000,
    "rate": 1000,
    "num_of_installment": 10
}'
```
- Get Loan detail
```
curl --location 'localhost:8080/loan/1'
```
- Get outstanding installments
```
curl --location 'localhost:8080/loan/1/outstandings'
```
- Get outstanding recap (Goal 1 and 2)
```
curl --location 'localhost:8080/loan/1/outstanding_recap'
```
- Pay an installment (Goal 3)
```
curl --location --request POST 'localhost:8080/loan/1/pay'
```
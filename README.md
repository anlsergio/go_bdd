# Golang BDD

This repository holds the result/practice of my study around implementing Acceptance Test Driven Development (ATDD) and Behavior Driven Development (BDD) by using Golang.

## Parking Lot API

This is a backend application designed to fullfil a request from a frontend squad of developers at the dummy's Southeast Airports Inc.

Nowadays Southeast Airports Inc. parking lot managers are having a hard time due to the lack of automation on parking lot cost calculation. 

So far, for every customer who drops their car at 5 different parking lot modalities, the parking lot staff need to do the maths manually, which takes time, let alone that a manual task like this is pretty error prone.

Here's the definition of the 5 different parking modalities how the cost is being calculated:

### Valet Parking Lot:

In Valet Parking Lot, the passenger drops his or her car off at the valet dropoff and gets a receipt to get the car back.

The Valet parking costs $18.00 a day, but for 5 hours or less there's a reduction of $6.00 as a discount.

<br />

### Short-Term Parking Lot:

For short-term parking, there are places for visitors dropping off or picking up other passengers.

Customers are charged $2.00 for the first hour only, then an additional $1.00 for each half hour passed the first hour. Although there's a daily maximum of $24.00.

It's important to emphasize that the $2.00 first hour is exclusive to the first hour only, and therefore, it doesn't apply for the 2nd day and on.

<br />

### Economy Parking Lot:

This parking lot is placed way apart from the airport, and that's what makes it cheaper for passengers. So that passengers need to make use of the bus shuttle which is offered at their disposal.

The cost calculation is as follow:

- $2.00 per hour
- there's a daily maximum of $9.00 until the next day
- after 24 hours (a day) the calculation sums up another $2.00 per hour up to the maximum of $9.00. And the same applies for any following day.
- there's also a weekly maximum of $54.00, which means that basically the 7th day is free of charge.

<br />

### Long-Term Parking Lot:

Both the Garage and the Surface modalities are variations of the Economy Parking Lot.

#### Long-Term - Garage Parking Lot:

For the long-term in the garage the costs are applied as follow:

- $2.00 per hour
- $12.00 as daily maximum
- 7th day is free of charge

<br />

#### Long-Term - Surface Parking Lot:

For the surface one, the cost calculation is as follow:

- $2.00 per hour
- $10.00 as daily maximum
- 7th day is free of charge

<br />

>Note: the cost calculation should consider the overall parking duration time, not the day specifically. In other words, 24 hours is condered a day, it doesn't matter if the check-in takes place one day and the check-out another day, it doesn't yeld one day plus another day.
Feature: Calculate price for Short-Term Parking
  As an app developer at Southeast Airports Inc.
  I want to be able to send a specific duration time in minutes to an API endpoint and receive the corresponding parking cost for Short-Term Parking
  So that I have a simple way to calculate the corresponding parking cost for an elapsed time spent at the parking lot, which can be used by the frontend apps

Scenario: Calculate Short-Term Parking for 30 minutes
  When a "GET" request is sent to the endpoint "/short-term/30"
  Then the HTTP response code should be 200
  And the response content should be "$2.00"

Scenario: Calculate Short-Term Parking for 1 hour
  When a "GET" request is sent to the endpoint "/short-term/60"
  Then the HTTP response code should be 200
  And the response content should be "$2.00"

Scenario: Calculate Short-Term Parking for 3 hours and 30 minutes
  When a "GET" request is sent to the endpoint "/short-term/210"
  Then the HTTP response code should be 200
  And the response content should be "$7.00"

Scenario: Calculate Short-Term Parking for 12 hours
  When a "GET" request is sent to the endpoint "/short-term/720"
  Then the HTTP response code should be 200
  And the response content should be "$24.00"

Scenario: Calculate Short-Term Parking for 12 hours and 30 minutes
  When a "GET" request is sent to the endpoint "/short-term/750"
  Then the HTTP response code should be 200
  And the response content should be "$24.00"

Scenario: Calculate Short-Term Parking for 1 day
  When a "GET" request is sent to the endpoint "/short-term/1440"
  Then the HTTP response code should be 200
  And the response content should be "$24.00"

Scenario: Calculate Short-Term Parking for 1 day and 30 minutes
  When a "GET" request is sent to the endpoint "/short-term/1470"
  Then the HTTP response code should be 200
  And the response content should be "$25.00"

Scenario: Calculate Short-Term Parking for 1 day and 1 hour
  When a "GET" request is sent to the endpoint "/short-term/1500"
  Then the HTTP response code should be 200
  And the response content should be "$26.00"

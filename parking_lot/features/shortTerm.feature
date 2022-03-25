Feature: Calculate price for Short-Term Parking
  As an app developer at Southeast Airports Inc.
  I want to be able to send a specific duration time in minutes to an API endpoint and receive the corresponding parking cost for Short-Term Parking
  So that I have a simple way to calculate the corresponding parking cost for an elapsed time spent at the parking lot, which can be used by the frontend apps

  Scenario Outline: Calculate Short-Term Parking
    When a "GET" request is sent to the endpoint "/short-term/<parking-duration>"
    Then the HTTP response code should be 200
    And the response content should be "<parking-cost>"
    Examples:
      | caption | parking-duration | parking-cost |
      | 30m     | 30               | $2.00        |
      | 1h      | 60               | $2.00        |
      | 3h30m   | 210              | $7.00        |
      | 12h     | 720              | $24.00       |
      | 12h30m  | 750              | $24.00       |
      | 1day    | 1440             | $24.00       |
      | 1day30m | 1470             | $25.00       |
      | 1day1h  | 1500             | $26.00       |

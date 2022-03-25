Feature: Calculate price for Valet Parking
  As an app developer at Southeast Airports Inc.
  I want to be able to send a specific amount of minutes to an API endpoint and receive the corresponding parking cost for Valed Parking
  So that I have a simple way to calculate the corresponding parking cost for an elapsed time spent at the parking lot, which can be used by the frontend apps

  Scenario Outline: Calculate Valet Parking Cost
    When a "GET" request is sent to the endpoint "/valet/<parking-duration>"
    Then the HTTP response code should be 200
    And the response content should be "<parking-cost>"
    Examples:
        | caption  | parking-duration | parking-cost |
        | 30m      | 30               | $12.00       |
        | 3h       | 180              | $12.00       |
        | 5h       | 300              | $12.00       |
        | 5h1m     | 301              | $18.00       |
        | 12h      | 720              | $18.00       |
        | 24h      | 1440             | $18.00       |
        | 1day1m   | 1441             | $36.00       |
        | 3days    | 4320             | $54.00       |
        | 1week    | 10080            | $126.00      |

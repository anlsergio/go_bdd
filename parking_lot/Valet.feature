Feature: Calculate price for Valet Parking
  As an app developer at Southeast Airports Inc.
  I want to be able to send a specific amount of minutes to an API endpoint and receive the corresponding parking cost for Valed Parking
  So that I have a simple way to calculate the corresponding parking cost for an elapsed time spent at the parking lot, which can be used by the frontend apps

  # Scenario: Calculate Valet Parking Cost for half an hour
  #   When I park my car in the Valet Parking Lot for 30 minutes
  #   Then I will have to pay $ 12.00

  Scenario: Calculate Valet Parking Cost for half an hour
    When a "GET" request is sent to the endpoint "/valet/30"
    Then the HTTP response code should be 200
    And the response content should be "$ 12.00"

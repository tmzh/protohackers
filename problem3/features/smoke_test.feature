Feature: Some Test
  TCP Clients can send and receive data 

  Scenario: Insert message
    Given the following message:
      "Hello world"
    And the messages are encoded into bytes
    When I send the bytes to TCP server
    Then I get the following response in bytes
      "Hello World"
    

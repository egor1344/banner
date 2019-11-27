# file: features/rest.feature

Feature: Rest server
  As Rest client of banner service

  Scenario: Add banner
    When I add banner rest-request to "http://rotation_banner"
    Then The json response add banner must contain status

  Scenario: Del banner
    When I del banner rest-request to "http://rotation_banner"
    Then The json response del banner must contain status

  Scenario: Count transition
    When I count transition banner rest-request to "http://rotation_banner"
    Then The json response  count transition must contain status

  Scenario: Get banner
    When I Get banner banner rest-request to "http://rotation_banner"
    Then The json response must contain id banner
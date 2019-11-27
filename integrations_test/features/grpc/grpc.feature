# file: features/server.feature

Feature: Grpc server
  As grpc client of banner service

  Scenario: Add banner
    When I add banner gprc-request to "rotation_banner"
    Then The response add banner must contain status

  Scenario: Del banner
    When I del banner gprc-request to "rotation_banner"
    Then The response del banner must contain status

  Scenario: Count transition
    When I count transition banner gprc-request to "rotation_banner"
    Then The response count transition must contain status

  Scenario: Get banner
    When I Get banner banner gprc-request to "rotation_banner"
    Then The response must contain id banner
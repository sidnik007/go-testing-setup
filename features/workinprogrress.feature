@wip
Feature: This is to ensure that the @wip tag is working


  Scenario: Checking the wip tag
    Given an unimplemented feature file
    When I add an wip tag to that feature file
    Then the feature file is ignored

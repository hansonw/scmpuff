Feature: branch command

  Background:
    Given a mocked home directory

  Scenario: Branch list is numbered
    Given I am in a git repository
    And a 1 byte file named "a.txt"
    And I successfully run the following commands:
      | git add a.txt |
      | git commit -m "init" |
    And I switch to git branch "branch_a"
    And I switch to existing git branch "master"
    And I switch to git branch "branch_b"
    And I switch to existing git branch "master"
    When I successfully run `scmpuff branch`
    Then the stdout from "scmpuff branch" should contain "* [1] master"
    And the stdout from "scmpuff branch" should contain "  [2] branch_a"
    And the stdout from "scmpuff branch" should contain "  [3] branch_b"

  Scenario: Detached HEAD is not numbered
    Given I am in a git repository
    And a 1 byte file named "a.txt"
    And I successfully run the following commands:
      | git add a.txt |
      | git commit -m "init" |
    And I successfully run `git checkout HEAD~0 --detach`
    When I successfully run `scmpuff branch`
    Then the stdout from "scmpuff branch" should contain "(HEAD detached"
    And the stdout from "scmpuff branch" should contain "  [1] master"


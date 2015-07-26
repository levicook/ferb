Feature: Front Matter

  Background:
    Given a successfully built app at "front-matter-app"
    When I cd into "build"

  Scenario: YAML Front Matter
    When I open "index.html"
    Then I should see:
    """
    This is the title
    """

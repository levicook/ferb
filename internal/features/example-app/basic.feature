Feature: building a website

  Background:
    Given a successfully built app at "example-app"
    When I cd into "build"

  Scenario: generates directory indexes
    Then the following files should exist:
       | about/index.html  |
       | blog-a/index.html |
       | blog-b/index.html |
       | index.html        |

  Scenario: blog a content
    Then the following files should exist:
        | blog-a/2014/01/14/hello-world/index.html   |
        | blog-a/2015/01/14/goodbye-world/index.html |
        | blog-a/2015/02/28/happy-v-day/index.html   |

  Scenario: blog a archives
    Then the following files should exist:
        | blog-a/2014/01/14/index.html  |
        | blog-a/2014/01/index.html     |
        | blog-a/2014/index.html        |
        | blog-a/2015/index.html        |
        | blog-a/2015/page/2/index.html |

  Scenario: blog b content
    Then the following files should exist:
       | blog-b/2014/01/15/hello-mars/index.html        |
       | blog-b/2014/01/19/happy-bday-mlk-jr/index.html |
       | blog-b/2015/01/15/goodbye-mars/index.html      |

  Scenario: blog b archives
    Then the following files should exist:
        | blog-b/2014/01/15/index.html     |
        | blog-b/2014/01/index.html        |
        | blog-b/2014/01/page/2/index.html |
        | blog-b/2014/index.html           |
        | blog-b/2014/page/2/index.html    |

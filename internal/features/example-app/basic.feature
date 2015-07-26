Feature: building a website

  Background:
    Given a successfully built app at "example-app"
    When I cd into "build"

  Scenario: directory indexes
    Then the following files should exist:
      | about/index.html  |
      | index.html        |

  Scenario: blog a indexes
    Then the following files should exist:
      | blog-a/index.html |
      | blog-a/page/2/index.html |
      | blog-a/page/3/index.html |
      | blog-a/page/4/index.html |
      | blog-a/page/5/index.html |
    And the following files should not exist:
      | blog-a/page/1/index.html |
      | blog-a/page/6/index.html |

  Scenario: blog a main content
    Then the following files should exist:
      | blog-a/2014/01/14/hello-world/index.html        |
      | blog-a/2015/01/14/goodbye-world/index.html      |
      | blog-a/2015/02/28/happy-v-day/index.html        |
      | blog-a/2015/11/11/happy-veterans-day/index.html |
      | blog-a/2015/11/26/happy-thanksgiving/index.html |

  Scenario: blog a chronological archives
    Then the following files should exist:
      | blog-a/2014/01/14/index.html  |
      | blog-a/2014/01/index.html     |
      | blog-a/2014/index.html        |
      | blog-a/2015/index.html        |
      | blog-a/2015/page/2/index.html |

  Scenario: blog b main content
    Then the following files should exist:
      | blog-b/2014/01/15/hello-mars/index.html        |
      | blog-b/2014/01/19/happy-bday-mlk-jr/index.html |
      | blog-b/2015/01/15/goodbye-mars/index.html      |

  Scenario: blog b chronological archives
    Then the following files should exist:
      | blog-b/2014/01/15/index.html     |
      | blog-b/2014/01/index.html        |
      | blog-b/2014/01/page/2/index.html |
      | blog-b/2014/index.html           |
      | blog-b/2014/page/2/index.html    |

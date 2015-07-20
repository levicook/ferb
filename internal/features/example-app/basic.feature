Feature: basic

  Scenario: building

    Given a successfully built app at "example-app"

    When I cd into "build"

    Then the following files should exist:
       | index.html                                 | 
       | about/index.html                           | 
       | blog-a/index.html                          | 
       | blog-a/2014/01/15/hello-world/index.html   | 
       | blog-a/2015/01/15/goodbye-world/index.html | 
       | blog-b/index.html                          | 
       | blog-b/2014/01/15/hello-mars/index.html    | 
       | blog-b/2015/01/15/goodbye-mars/index.html  | 

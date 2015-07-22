
Ferb
====

Nothing to see here, move along :)

Ferb is a static site generator.
Ferb is besties with git and s3.

```
  source/                              ------>  dist/
  .  content/                          ------>  .  content/
  .  .  sitemap.xml.tmpl               ------>  .  .  sitemap.xml
  .  .  index.html.tmpl                ------>  .  .  index.html
  .  .  blerg/                         ------>  .  .  blerg/
  .  .  .  feed.xml.tmpl               ------>  .  .  .  feed.xml
  .  .  .  index.html.tmpl             ------>  .  .  .  index.html
  .  .  .  2015-05-13-hello-world.md   ------>  .  .  .  2015/05/13/hello-world/index.html
  .  .  .  2015-05-13-bright-ideas.md  ------>  .  .  .  2015/05/13/bright-ideas/index.html
  .  .  .  2015-05-14-insights.md      ------>  .  .  .  2015/05/14/insights/index.html
  .  .  .  2015-05-15-ideas.md         ------>  .  .  .  2015/05/15/ideas/index.html
  .  .  .  2015-06-15-observations.md  ------>  .  .  .  2015/06/15/observations/index.html
  .  .  .  2015-07-15-thoughts.md      ------>  .  .  .  2015/07/15/thoughts/index.html
  .  .  .                                       .  .  .
  .  .  .                              ------>  .  .  .  2015/index.html               | page 1 of 6 \
  .  .  .                              ------>  .  .  .  2015/page/2/index.html        | page 2 of 6  | 2015 Archive (1 article/page)
  .  .  .                              ------>  .  .  .  2015/page/3/index.html        | page 3 of 6  |
  .  .  .                              ------>  .  .  .  2015/page/4/index.html        | page 4 of 6  |
  .  .  .                              ------>  .  .  .  2015/page/5/index.html        | page 5 of 6  |
  .  .  .                              ------>  .  .  .  2015/page/6/index.html        | page 6 of 6 /  
  .  .  .                                       .  .  .
  .  .  .                              ------>  .  .  .  2015/05/index.html            | page 1 of 3 \
  .  .  .                              ------>  .  .  .  2015/05/page/2/index.html     | page 2 of 3  | May 2015 Archive (1 article/page)
  .  .  .                              ------>  .  .  .  2015/05/page/3/index.html     | page 3 of 3 /
  .  .  .                                       .  .  .
  .  .  .                              ------>  .  .  .  2015/05/13/index.html         | page 1 of 2 \ May 13th, 2015 Archive (1 article/page)
  .  .  .                              ------>  .  .  .  2015/05/13/page/2/index.html  | page 2 of 2 / 
  .  .  .                                       .  .  .
  .  .  .                              ------>  .  .  .  2015/05/14/index.html         | page 1 of 1 - May 14th, 2015 Archive
  .  .  .                              ------>  .  .  .  2015/05/15/index.html         | page 1 of 1 - May 15th, 2015 Archive
  .  .  .                                       .  .  .
  .  .  .                              ------>  .  .  .  2015/06/index.html            | page 1 of 1 - June 2015 Archive
  .  .  .                              ------>  .  .  .  2015/06/15/index.html         | page 1 of 1 - June 15th, 2015 Archive
  .  .  .                              ------>  .  .  .  2015/07/index.html            | page 1 of 1 - July 2015 Archive
  .  .  .                              ------>  .  .  .  2015/07/15/index.html         | page 1 of 1 - July 15th, 2015 Archive
  .  .  .                                       .  .  .
  .  .  images/                        ------>  .  .  images/
  .  .  .  favicon-16x16.png           ------>  .  .  .  favicon-16x16-711a8021.png
  .  .  .  favicon-32x32.png           ------>  .  .  .  favicon-32x32-3eb25150.png
  .  .  .  social/                     ------>  .  .  .  social/
  .  .  .  .  facebook.png             ------>  .  .  .  .  facebook-b7428b86.png
  .  .  .  .  linkedin.png             ------>  .  .  .  .  linkedin-739493a2.png
  .  .  .  .  twitter.png              ------>  .  .  .  .  twitter-73bab5a2.png
  .  .  .  .  youtube.png              ------>  .  .  .  .  youtube-fb5dd4c9.png
  .  .  javascripts/                   ------>  .  .  javascripts/
  .  .  .  site.js                     ------>  .  .  .  site-3eb25150.js
  .  .  stylesheets/                   ------>  .  .  stylesheets/
  .  .  .  site.css                    ------>  .  .  .  site-711a8021.css
  .  layouts/
  .  partials/

```

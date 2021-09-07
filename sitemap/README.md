# Sitemap Builder

## Details

A sitemap is basically a map of all of the pages within a specific domain. They are used by search engines and other tools to inform them of all of the pages on your domain.

Our application will accept a url from the command line and come up with a sitemap rooted at the provieded url.
The sitemap builder will then output the data in the following XML format:

```xml
<?xml version="1.0" encoding="UTF-8"?>
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
  <url>
    <loc>http://www.example.com/</loc>
  </url>
  <url>
    <loc>http://www.example.com/dogs</loc>
  </url>
</urlset>
```

This implementation follows the [standard sitemap protocol](https://www.sitemaps.org/index.html)\*

Where each page is listed in its own `<url>` tag and includes the `<loc>` tag inside of it.

In order to complete this task we will be making use of the [link parser](https://github.com/Smelton01/go-basics/tree/master/link) in order to parse HTML pages for links.

Links going to a different domain will be excluded from the final sitemap.

### Notes

**1. Cyclical links will be avoided.**

That is, page `abc.com` may link to page `abc.com/about`, and then the about page may link back to the home page (`abc.com`). These cycles can also occur over many pages, for instance you might have:

```
/about -> /contact
/contact -> /pricing
/pricing -> /testimonials
/testimonials -> /about
```

Where the cycle takes 4 links to finally reach it, but there is indeed a cycle.

This is important to remember because you don't want your program to get into an infinite loop where it keeps visiting the same few pages over and over. If you are having issues with this, the bonus exercise might help temporarily alleviate the problem but we will cover how to avoid this entirely in the screencasts for this exercise.

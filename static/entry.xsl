<xsl:stylesheet
    exclude-result-prefixes="#all"
    expand-text="yes"
    version="3.0"
    xmlns:xsl="http://www.w3.org/1999/XSL/Transform"
    xmlns:atom="http://www.w3.org/2005/Atom"
    xmlns:comet="https://comet.pub/Atom"
    xmlns="http://www.w3.org/1999/xhtml"
    xmlns:svg="http://www.w3.org/2000/svg"
    xmlns:xlink="http://www.w3.org/1999/xlink">
    <xsl:output method="html" encoding="utf-8" version="5"/>
    <xsl:include href="cards.xsl"/>
    <xsl:template match="/">
        <xsl:variable name="entry" select="atom:feed/atom:entry[1]"/>
        <html lang="en" data-theme="light">
            <head>
                <!-- Common -->
                <meta charset="UTF-8"/>
                <meta name="viewport" content="width=device-width, initial-scale=1.0"/>
                <title>
                    <xsl:value-of select="atom:feed/atom:title"/>
                </title>
                <!-- Favicons -->
                <link rel="icon">
                    <xsl:attribute name="href">
                        <xsl:value-of select="/atom:feed/atom:icon"/>
                    </xsl:attribute>
                </link>
                <!-- Comet styles -->
                <link rel="stylesheet" href="/static/css/feed.css"/>
                <!-- HTMX and local scripts -->
                <script src="/static/js/feed.js"></script>
            </head>
            <body>
                <xsl:apply-templates select="atom:feed" />
                <main>
                    <xsl:choose>
                        <xsl:when test="$entry/atom:content != ''">
                            <section class="h-entry">
                                <header>
                                    <h1 class="p-name">
                                        <xsl:value-of select="$entry/atom:title"/>
                                    </h1>
                                    <xsl:apply-templates select="$entry/atom:published"/>
                                </header>
                                <hr/>
                                <blockquote class="p-summary">
                                    <xsl:value-of select="$entry/atom:summary"/>
                                </blockquote>
                                <div class="p-content">
                                    <xsl:apply-templates select="$entry/atom:content"/>
                                </div>
                            </section>
                        </xsl:when>
                        <xsl:otherwise>
                            <xsl:apply-templates select="$entry" />
                        </xsl:otherwise>
                    </xsl:choose>
                </main>
            </body>
        </html>
    </xsl:template>
    
    <xsl:template match="atom:feed">
        <header>
            <a href="/">
                <img width="48" height="48" aria-hidden="true">
                    <xsl:attribute name="src">
                        <xsl:value-of select="atom:icon"/>
                    </xsl:attribute>
                </img>
                <h1 class="sr-only">
                    <xsl:value-of select="atom:title"/>
                </h1>
            </a>
            <nav>
                <a href="/articles">Articles</a>
                <a href="/notes">Notes</a>
            </nav>
            <theme-toggle></theme-toggle>
            <a href="/atom.xml" class="icon-rss">
                <span class="sr-only">RSS Feed</span>
            </a>
        </header>
    </xsl:template>
    
    <xsl:template match="atom:author">
        <div class="h-card p-author">
            <img class="u-logo">
                <xsl:attribute name="src">
                    <xsl:value-of select="atom:author/comet:avatar" />
                </xsl:attribute>
                <xsl:attribute name="alt">
                    <xsl:value-of select="atom:author/atom:name" />
                </xsl:attribute>
            </img>
            <p class="p-name">
                <b><xsl:value-of select="atom:author/comet:username" /></b>
            </p>
        </div>
    </xsl:template>
    
    <xsl:template match="atom:content|atom:summary">
        <xsl:value-of select="." disable-output-escaping="yes"></xsl:value-of>
    </xsl:template>
    
    <xsl:template match="atom:published">
        <locale-date class="dt-published">
            <xsl:attribute name="datetime">
                <xsl:value-of select="."/>
            </xsl:attribute>
            <xsl:value-of select="substring(., 0, 11)" />
        </locale-date>
    </xsl:template>
    
    <xsl:template match="atom:category">
        <li>
            <a rel="category tag" role="button" class="p-category tag secondary">
                <xsl:attribute name="href">
                    <xsl:value-of select="atom:id" />
                </xsl:attribute>
                #
                <xsl:value-of select="." />
            </a>
        </li>
    </xsl:template>
</xsl:stylesheet>
<xsl:stylesheet
    exclude-result-prefixes="#all"
    expand-text="yes"
    version="3.0"
    xmlns:xsl="http://www.w3.org/1999/XSL/Transform"
    xmlns:atom="http://www.w3.org/2005/Atom"
    xmlns:comet="https://comet.pub/Atom"
    xmlns="http://www.w3.org/1999/xhtml"
    xmlns:xlink="http://www.w3.org/1999/xlink">
    <xsl:output method="html" encoding="utf-8" indent="yes"/>
    <xsl:include href="cards.xsl"/>
    
    <xsl:template match="/">
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
                <!-- Comet styles and scripts -->
                <link rel="stylesheet" href="/static/css/feed.css" />
                <script src="/static/js/feed.js" type="module"></script>
            </head>
            <body>
                <xsl:apply-templates select="atom:feed" />
                <main>
                    <xsl:apply-templates select="atom:feed/atom:category" />
                    <ol role="list">
                        <xsl:for-each select="atom:feed/atom:entry">
                            <li>
                                <xsl:apply-templates select="." />
                            </li>
                        </xsl:for-each>
                        
                    </ol>
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
                <h1>
                    <xsl:value-of select="atom:title"/>
                </h1>
            </a>
            <theme-toggle></theme-toggle>
            <a href="/atom.xml" class="icon-rss">
                <span class="sr-only">RSS Feed</span>
            </a>
        </header>
    </xsl:template>
    
    <xsl:template match="atom:feed/atom:category">
        <h2 class="p-summary">
            <span class="sr-only">Posts tagged</span>
            <xsl:value-of select="@label" />
        </h2>
    </xsl:template>
</xsl:stylesheet>
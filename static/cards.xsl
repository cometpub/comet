<xsl:stylesheet
    exclude-result-prefixes="#all"
    expand-text="yes"
    version="3.0"
    xmlns:xsl="http://www.w3.org/1999/XSL/Transform"
    xmlns:atom="http://www.w3.org/2005/Atom"
    xmlns:comet="https://comet.pub/Atom"
    xmlns="http://www.w3.org/1999/xhtml"
    xmlns:xlink="http://www.w3.org/1999/xlink">
    <xsl:template match="atom:entry">
        <article class="h-entry">
            <header>
                <xsl:apply-templates select="atom:author"/>
                <xsl:apply-templates select="atom:published"/>
            </header>
            <xsl:if test="atom:content != '' and atom:title != ''">
                <h2 class="p-name">
                    <xsl:value-of select="atom:title" />
                </h2>
            </xsl:if>
            <xsl:if test="not(atom:content) and atom:link[starts-with(@type, 'image/')]">
                <xsl:choose>
                    <xsl:when test="count(atom:link[starts-with(@type, 'image/')]) = 1">
                        <xsl:apply-templates select="atom:link[starts-with(@type, 'image/')]"/>
                    </xsl:when>
                    <xsl:otherwise>
                        <ul role="list">
                            <xsl:apply-templates select="atom:link[starts-with(@type, 'image/')]"/>
                        </ul>
                    </xsl:otherwise>
                </xsl:choose>
            </xsl:if>
            <xsl:if test="atom:summary != ''">
                <div class="p-summary">
                    <xsl:apply-templates select="atom:summary"/>
                </div>
            </xsl:if>
            <xsl:if test="atom:category != ''">
                <ul role="list">
                    <xsl:apply-templates select="atom:category" />
                </ul>
            </xsl:if>
            <footer>
                <a class="u-url">
                    <xsl:attribute name="href">
                        <xsl:value-of select="atom:link[@rel='self']/@href" />
                    </xsl:attribute>
                    Full Post
                </a>
                <button aria-label="Share a direct link to this post" class="icon-share">
                </button>
            </footer>
        </article>
    </xsl:template>
    
    <xsl:template match="atom:entry/atom:author">
        <div class="h-card p-author">
            <img class="u-logo">
                <xsl:attribute name="src">
                    <xsl:value-of select="comet:avatar" />
                </xsl:attribute>
                <xsl:attribute name="alt">
                    <xsl:value-of select="atom:name" />
                </xsl:attribute>
            </img>
            <p class="p-name">
                <b><xsl:value-of select="comet:username" /></b>
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
            <a role="button" class="p-category" rel="category tag">
                <xsl:attribute name="href">
                    <xsl:value-of select="concat('/c/', .)" />
                </xsl:attribute>
                #
                <xsl:value-of select="." />
            </a>
        </li>
    </xsl:template>
    
    <xsl:template match="atom:link[starts-with(@type, 'image/')]">
        <img alt="" class="u-photo">
            <xsl:attribute name="src">
                <xsl:value-of select="concat(@href, '?thumb=900x0')" />
            </xsl:attribute>
            <xsl:attribute name="width">
                <xsl:value-of select="@comet:width" />
            </xsl:attribute>
            <xsl:attribute name="height">
                <xsl:value-of select="@comet:height" />
            </xsl:attribute>
        </img>
    </xsl:template>
</xsl:stylesheet>
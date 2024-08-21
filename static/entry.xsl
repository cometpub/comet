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
                <style>
                    @layer base, components, utilities, view;
                    
                    /* src/css/base/reset.css */
                    @layer base.reset {
                      *,
                      *::before,
                      *::after {
                        box-sizing: border-box;
                      }
                      * {
                        margin: 0;
                      }
                      body {
                        line-height: 1.8;
                        -webkit-font-smoothing: antialiased;
                      }
                      img,
                      picture,
                      video,
                      canvas,
                      svg {
                        display: block;
                        max-width: 100%;
                      }
                      input,
                      button,
                      textarea,
                      select {
                        font: inherit;
                        background-color: var(--surface-2);
                        color: inherit;
                      }
                      p,
                      h1,
                      h2,
                      h3,
                      h4,
                      h5,
                      h6 {
                        overflow-wrap: break-word;
                        display: block;
                      }
                      blockquote {
                        border-inline-start: 0.25rem solid var(--text-accent);
                        padding: var(--space-s);
                        padding-inline-end: 0;
                      }
                      blockquote > p {
                        display: initial;
                      }
                      h1,
                      h2,
                      h3,
                      h4,
                      h5,
                      h6 {
                        line-height: 1.2;
                      }
                      hr {
                        border-color: var(--surface-3);
                        margin-block: var(--space-m-l);
                      }
                      fieldset {
                        border: none;
                        padding: 0;
                      }
                      textarea {
                        width: 100%;
                        resize: vertical;
                      }
                      :where(ul, ol, dl)[role=list] {
                        margin: 0;
                        padding: 0;
                      }
                      :where(ul, ol, dl)[role=list] > li {
                        list-style: none;
                      }
                      ul {
                        list-style-position: inside;
                        list-style-type: disc;
                      }
                      ul > li > p:first-child {
                        display: initial;
                      }
                      li:has(> p:first-child > input[type=checkbox]:first-child) {
                        list-style: none;
                      }
                    }
                    
                    /* src/css/base/prose.css */
                    @layer base.prose {
                      :where(.p-summary, .p-content) > * + * {
                        margin-block-start: var(--space-s);
                      }
                      :where(.p-summary, .p-content) :where(h2, h3, h4, h5, h6) {
                        margin: calc(1rem + 0.75em) 0 calc(0.5rem + 0.15em);
                      }
                      :where(.p-summary, .p-content) ul {
                        padding-inline-start: 0;
                      }
                    }
                    
                    /* src/css/base/theme.css */
                    @layer base.theme {
                      :root {
                        --neutral-50: #f8f9fa;
                        --neutral-100: #f1f3f5;
                        --neutral-200: #e9ecef;
                        --neutral-300: #dee2e6;
                        --neutral-400: #ced4da;
                        --neutral-500: #adb5bd;
                        --neutral-600: #868e96;
                        --neutral-700: #495057;
                        --neutral-800: #343a40;
                        --neutral-900: #212529;
                        --neutral-1000: #16191d;
                        --neutral-1100: #0d0f12;
                        --neutral-1200: #030507;
                        --primary-50: #e3fafc;
                        --primary-100: #c5f6fa;
                        --primary-200: #99e9f2;
                        --primary-300: #66d9e8;
                        --primary-400: #3bc9db;
                        --primary-500: #22b8cf;
                        --primary-600: #15aabf;
                        --primary-700: #1098ad;
                        --primary-800: #0c8599;
                        --primary-900: #0b7285;
                        --primary-1000: #095c6b;
                        --primary-1100: #074652;
                        --primary-1200: #053038;
                        --surface-1: var(--neutral-200);
                        --surface-2: var(--neutral-50);
                        --surface-3: var(--neutral-300);
                        --surface-primary: var(--primary-900);
                        --text-primary: var(--neutral-1000);
                        --text-secondary: var(--neutral-600);
                        --text-accent: var(--primary-900);
                        --text-alt: var(--neutral-50);
                        --color-bg: var(--neutral-50);
                        --color-fg: var(--neutral-1100);
                        --color-link: var(--primary-900);
                        --font-serif:
                          Rockwell,
                          "Rockwell Nova",
                          "Roboto Slab",
                          "DejaVu Serif",
                          "Sitka Small",
                          serif;
                        font-weight: normal;
                        --font-sans:
                          Seravek,
                          "Gill Sans Nova",
                          Ubuntu,
                          Calibri,
                          "DejaVu Sans",
                          source-sans-pro,
                          sans-serif;
                        --font-mono:
                          ui-monospace,
                          "Cascadia Code",
                          "Source Code Pro",
                          Menlo,
                          Consolas,
                          "DejaVu Sans Mono",
                          monospace;
                        --font-size-s: clamp(0.8333rem, 0.8101rem + 0.1159vi, 0.9rem);
                        --font-size-m: clamp(1rem, 0.9565rem + 0.2174vi, 1.125rem);
                        --font-size-l: clamp(1.2rem, 1.1283rem + 0.3587vi, 1.4063rem);
                        --font-size-xl: clamp(1.44rem, 1.3295rem + 0.5527vi, 1.7578rem);
                        --font-size-2xl: clamp(1.728rem, 1.5648rem + 0.8161vi, 2.1973rem);
                        --font-size-3xl: clamp(2.0736rem, 1.8395rem + 1.1704vi, 2.7466rem);
                        --font-size-4xl: clamp(2.4883rem, 2.1597rem + 1.6433vi, 3.4332rem);
                        --space-3xs: clamp(0.25rem, 0.2283rem + 0.1087vi, 0.3125rem);
                        --space-2xs: clamp(0.5rem, 0.4783rem + 0.1087vi, 0.5625rem);
                        --space-xs: clamp(0.75rem, 0.7065rem + 0.2174vi, 0.875rem);
                        --space-s: clamp(1rem, 0.9565rem + 0.2174vi, 1.125rem);
                        --space-m: clamp(1.5rem, 1.4348rem + 0.3261vi, 1.6875rem);
                        --space-l: clamp(2rem, 1.913rem + 0.4348vi, 2.25rem);
                        --space-xl: clamp(3rem, 2.8696rem + 0.6522vi, 3.375rem);
                        --space-2xl: clamp(4rem, 3.8261rem + 0.8696vi, 4.5rem);
                        --space-3xl: clamp(6rem, 5.7391rem + 1.3043vi, 6.75rem);
                        --space-xs-m: clamp(0.75rem, 0.4239rem + 1.6304vi, 1.6875rem);
                        --space-s-m: clamp(1rem, 0.7609rem + 1.1957vi, 1.6875rem);
                        --space-m-l: clamp(1.5rem, 1.2391rem + 1.3043vi, 2.25rem);
                      }
                      [data-theme=dark] {
                        --surface-1: var(--neutral-900);
                        --surface-2: var(--neutral-800);
                        --surface-3: var(--neutral-700);
                        --surface-primary: var(--primary-300);
                        --text-primary: var(--neutral-50);
                        --text-secondary: var(--neutral-300);
                        --text-accent: var(--primary-300);
                        --text-alt: var(--neutral-800);
                      }
                      html {
                        font-family: var(--font-sans);
                        font-size: var(--font-size-m);
                        max-width: 80ch;
                        margin: auto;
                        background-color: var(--surface-1);
                        color: var(--text-primary);
                        accent-color: var(--surface-primary);
                      }
                      body {
                        padding: var(--space-xs-m);
                      }
                      main {
                        margin-block: var(--space-m-l);
                        display: flex;
                        flex-direction: column;
                        gap: var(--space-m-l);
                      }
                      html:has(main > article:only-child) {
                        height: 100%;
                      }
                      html:has(main > article:only-child) body {
                        height: 100%;
                        display: flex;
                        flex-direction: column;
                      }
                      html:has(main > article:only-child) main {
                        display: flex;
                        align-items: center;
                        justify-content: center;
                        flex: 1;
                      }
                      h1,
                      h2,
                      h3,
                      h4,
                      h5,
                      h6 {
                        text-wrap: balance;
                        text-wrap: pretty;
                        font-family: var(--font-sans);
                      }
                      h1 {
                        font-size: var(--font-size-4xl);
                      }
                      h2 {
                        font-size: var(--font-size-3xl);
                      }
                      h3 {
                        font-size: var(--font-size-2xl);
                      }
                      h4 {
                        font-size: var(--font-size-xl);
                      }
                      h5 {
                        font-size: var(--font-size-l);
                      }
                      h6 {
                        font-size: var(--font-size-l);
                      }
                      small {
                        font-size: var(--font-size-s);
                      }
                      code,
                      pre,
                      locale-date {
                        font-family: var(--font-mono);
                      }
                      :is(a, a:visited) {
                        color: var(--color-link);
                      }
                      :is(a, a:visited):where(:hover, :focus-visible) {
                        opacity: 0.7;
                      }
                    }
                    
                    /* src/css/components/components.css */
                    @layer components {
                      body > header {
                        display: flex;
                        align-items: center;
                        gap: var(--space-xs);
                      }
                      body > header nav {
                        display: flex;
                        gap: var(--space-xs);
                      }
                      body > header nav a {
                        color: inherit;
                        text-decoration: none;
                        font-size: var(--font-size-m);
                      }
                      body > header theme-toggle {
                        margin-inline-start: auto;
                      }
                      :where([role=button], [type=button], [type=file]::file-selector-button, [type=reset], [type=submit], button) {
                        --color: var(--text-alt);
                        --background: var(--surface-primary);
                        --border-color: var(--background);
                        --font-size: var(--font-size-m);
                        font-size: var(--font-size);
                        color: var(--color);
                        background-color: var(--background);
                        border: 2px solid var(--border-color);
                        border-radius: 0.5em;
                        padding: 0.125em 0.67em;
                        text-decoration: none;
                        cursor: pointer;
                      }
                      :where([role=button], [type=button], [type=file]::file-selector-button, [type=reset], [type=submit], button).outline {
                        --border-color: var(--text-accent);
                      }
                      :where([role=button], [type=button], [type=file]::file-selector-button, [type=reset], [type=submit], button):where(.outline, .ghost, [class^=icon-]:empty, [class*=" icon-"]:empty) {
                        --background: transparent;
                        --color: var(--text-accent);
                      }
                      :where([role=button], [type=button], [type=file]::file-selector-button, [type=reset], [type=submit], button):where([class^=icon-]:empty, [class*=" icon-"]:empty) {
                        font-size: var(--font-size-l);
                      }
                      :where([role=button], [type=button], [type=file]::file-selector-button, [type=reset], [type=submit], button)[role=link] {
                        --color: var(--text-accent);
                        --background: transparent;
                        border: none;
                        padding: 0;
                      }
                    }
                    
                    /* src/css/fonts/icons.css */
                    @layer components {
                      @font-face {
                        font-family: "icomoon";
                        src: url(/static/fonts/icons/icomoon.eot?zbf2tp);
                        src:
                          url(/static/fonts/icons/icomoon.eot?zbf2tp#iefix) format("embedded-opentype"),
                          url(/static/fonts/icons/icomoon.ttf?zbf2tp) format("truetype"),
                          url(/static/fonts/icons/icomoon.woff?zbf2tp) format("woff"),
                          url(/static/fonts/icons/icomoon.svg?zbf2tp#icomoon) format("svg");
                        font-weight: normal;
                        font-style: normal;
                        font-display: block;
                      }
                      [class^=icon-],
                      [class*=" icon-"] {
                        font-family: "icomoon" !important;
                        speak: never;
                        font-style: normal;
                        font-weight: normal;
                        font-variant: normal;
                        text-transform: none;
                        line-height: 1;
                        -webkit-font-smoothing: antialiased;
                        -moz-osx-font-smoothing: grayscale;
                      }
                      .icon-wb_sunny:before {
                        content: "\e900";
                      }
                      .icon-brightness:before {
                        content: "\e901";
                      }
                      .icon-quotes-right:before {
                        content: "\e978";
                      }
                      .icon-lock:before {
                        content: "\e98f";
                      }
                      .icon-unlocked:before {
                        content: "\e990";
                      }
                      .icon-attachment:before {
                        content: "\e9cd";
                      }
                      .icon-bold:before {
                        content: "\ea62";
                      }
                      .icon-italic:before {
                        content: "\ea64";
                      }
                      .icon-share:before {
                        content: "\ea7d";
                      }
                      .icon-rss:before {
                        content: "\ea9b";
                      }
                      .icon-image:before {
                        content: "\e90d";
                      }
                      .icon-images:before {
                        content: "\e90e";
                      }
                      .icon-cancel-circle:before {
                        content: "\ea0d";
                      }
                    }
                    
                    /* src/css/utils/utilities.css */
                    @layer utilities {
                      .sr-only {
                        border: 0;
                        padding: 0;
                        margin: 0;
                        position: absolute !important;
                        height: 1px;
                        width: 1px;
                        overflow: hidden;
                        clip: rect(1px 1px 1px 1px);
                        clip: rect(1px, 1px, 1px, 1px);
                        clip-path: inset(50%);
                        white-space: nowrap;
                      }
                      .sr-only.sr-only-focusable:focus {
                        width: auto;
                        height: auto;
                        padding: 0;
                        margin: 0;
                        overflow: visible;
                        clip: auto;
                        clip-path: initial;
                        white-space: normal;
                      }
                    }
                    
                    /* src/css/components/cards.css */
                    @layer view {
                      a[rel~=category] {
                        --font-size: var(--font-size-s);
                        font-weight: 700;
                      }
                      ul:has(a[rel~=category]) {
                        display: flex;
                        gap: var(--space-xs);
                        flex-wrap: wrap;
                      }
                      .h-entry .p-author {
                        display: flex;
                        align-items: center;
                        gap: var(--space-xs);
                      }
                      .h-entry .p-author .u-logo {
                        border-radius: 9999px;
                        height: var(--space-xl);
                        width: var(--space-xl);
                        object-fit: cover;
                      }
                      .h-entry:not(:has(.p-content)) {
                        background-color: var(--surface-2);
                        max-width: 50ch;
                        width: 100%;
                        border-radius: 0.33em;
                        display: flex;
                        flex-direction: column;
                        gap: var(--space-s);
                        padding: var(--space-s);
                      }
                      .h-entry:not(:has(.p-content)) .p-author + * {
                        margin-inline-start: auto;
                      }
                      .h-entry:not(:has(.p-content)) .p-author .p-name {
                        font-size: var(--font-size-m);
                      }
                      .h-entry:not(:has(.p-content)) .p-name {
                        font-size: var(--font-size-xl);
                      }
                      .h-entry:not(:has(.p-content)) :where(header, footer) {
                        display: flex;
                        align-items: center;
                        justify-content: space-between;
                        gap: var(--space-xs);
                      }
                      .h-entry:has(.p-content) header {
                        display: grid;
                        gap: var(--space-s);
                        align-items: center;
                        grid-template-columns: 1fr auto;
                        grid-template-areas: "categories categories" "title title" "author published";
                      }
                      .h-entry:has(.p-content) header .p-name {
                        grid-area: title;
                      }
                      .h-entry:has(.p-content) header .dt-published {
                        grid-area: published;
                      }
                      .h-entry:has(.p-content) header .p-author {
                        grid-area: author;
                      }
                      .h-entry:has(.p-content) header ul:has([rel~=category]) {
                        grid-area: categories;
                      }
                      .h-entry:has(.p-content) footer .u-url {
                        margin-inline-end: -0.67em;
                      }
                      .h-entry:has(.p-content) .p-summary {
                        margin-block: var(--space-m);
                        font-size: var(--font-size-l);
                        line-height: 1.5;
                      }
                      .u-photo {
                        height: auto;
                      }
                    }
                    
                    /* src/css/feed.css */
                    @layer view {
                      ol:has(.h-entry) {
                        display: flex;
                        flex-direction: column;
                        gap: var(--space-s-m);
                        margin-inline: auto;
                      }
                      h2.p-summary {
                        text-align: center;
                      }
                      nav[aria-label=Pagination] {
                        display: grid;
                        grid-template-columns: auto auto 1fr auto auto;
                        grid-template-areas: "first previous . next last";
                        gap: var(--space-s);
                        width: 100%;
                        max-width: 50ch;
                        margin-inline: auto;
                      }
                      nav[aria-label=Pagination] a[rel=first] {
                        grid-area: first;
                      }
                      nav[aria-label=Pagination] a[rel=previous] {
                        grid-area: previous;
                      }
                      nav[aria-label=Pagination] a[rel=next] {
                        grid-area: next;
                      }
                      nav[aria-label=Pagination] a[rel=last] {
                        grid-area: last;
                      }
                      form[action="/publish"] {
                        padding: var(--space-s-m);
                        border-radius: 0.5em;
                        background-color: var(--surface-2);
                        width: 100%;
                        max-width: 50ch;
                        margin-inline: auto;
                        display: flex;
                        flex-direction: column;
                        gap: var(--space-s);
                      }
                    }
                </style>
                <script src="/static/js/feed.js" type="module"></script>
            </head>
            <body>
                <xsl:variable name="entry" select="atom:feed/atom:entry[1]"/>
                <xsl:apply-templates select="atom:feed" />
                <main>
                    <xsl:choose>
                        <xsl:when test="$entry/atom:content != ''">
                            <section class="h-entry">
                                <header>
                                    <xsl:if test="$entry/atom:category != ''">
                                        <ul role="list">
                                            <xsl:apply-templates select="$entry/atom:category" />
                                        </ul>
                                    </xsl:if>
                                    <h1 class="p-name">
                                        <xsl:value-of select="$entry/atom:title"/>
                                    </h1>
                                    <xsl:apply-templates select="$entry/atom:author" />
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
                    <script>
                        // https://bugzilla.mozilla.org/show_bug.cgi?id=98168
document.querySelectorAll(".p-content, .p-summary").forEach(function (elem) {
  if (elem.firstChild?.nodeName == "#text") {
    elem.innerHTML = elem.innerText;
  }
});
                    </script>
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
    </xsl:template>    <xsl:template match="atom:entry">
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
                <share-button>
                    <xsl:attribute name="title">
                        <xsl:value-of select="atom:title" />
                    </xsl:attribute>
                    <xsl:attribute name="title">
                        <xsl:value-of select="atom:title" />
                    </xsl:attribute>
                    <xsl:attribute name="url">
                        <xsl:value-of select="atom:link[@rel='self']/@href" />
                    </xsl:attribute>
                    <span aria-label="Share a direct link to this post" class="icon-share"></span>
                </share-button>
            </footer>
        </article>
    </xsl:template>
    
    <xsl:template match="atom:entry/atom:author">
        <div class="h-card p-author">
            <img class="u-logo" decoding="async" loading="lazy">
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
                    <xsl:value-of select="concat('/category/', .)" />
                </xsl:attribute>
                #
                <xsl:value-of select="." />
            </a>
        </li>
    </xsl:template>
    
    <xsl:template match="atom:link[starts-with(@type, 'image/')]">
        <img alt="" class="u-photo" decoding="async" loading="lazy">
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
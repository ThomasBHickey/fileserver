<?xml version="1.0" encoding="utf-8"?>
<xsl:stylesheet xmlns:xsl="http://www.w3.org/1999/XSL/Transform"
  xmlns:xr = "http://viaf.org/ontology/xr/#"
  version="1.0">

  <xsl:output method="html" version="4.1" standalone="yes" indent="yes" doctype-public="-//W3C//DTD XHTML 1.0 Strict//EN" doctype-system="http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd" encoding="UTF-8" />

<xsl:template match="/">
    <html xmlns="http://www.w3.org/1999/xhtml" xml:lang="en" lang="en" > 
      <xsl:call-template name="editHead"/>
	  <xsl:call-template name="editBody"/>
    </html>
  </xsl:template>

  <xsl:template name="editHead">
	<head>
		<title>ACE in Action</title>
		<style type="text/css" media="screen">
    #editor { 
        position: absolute;
        top: 0;
        right: 0;
        bottom: 0;
        left: 0;
    }
		</style>
		<script src="ace-builds/src-min-noconflict/ace.js" type="text/javascript" charset="utf-8"></script>	</head>
  </xsl:template>

  <xsl:template name="editBody">
	<xsl:variable name="content"><xsl:value-of select="//content"/></xsl:variable>
	<body>
	<div id="editor"><xsl:value-of select="$content"/></div>
    

<script>
    var editor = ace.edit("editor");
    editor.setTheme("ace/theme/monokai");
    editor.getSession().setMode("ace/mode/javascript");
</script>
</body>
</xsl:template>
</xsl:stylesheet>

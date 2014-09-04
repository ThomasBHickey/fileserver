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
		<!--style type="text/css" media="screen">
    #editor { 
        position: absolute;
        top: 0;
        right: 0;
        bottom: 0;
        left: 0;
    }
		</style-->
<style type="text/css" media="screen">
    textarea {
        display: block;
        margin: auto;
        width: 600px;
        height: 5px;
    }
    #editor {
        display: block;
        margin: auto;
        width: 600px;
        height: 300px;
        border: 1px solid #888;
    }
    form {
        text-align: center;
    }
</style>
	<script src="ace-builds/src-min-noconflict/ace.js" type="text/javascript" charset="utf-8"></script>
	<script src="jquery-2.1.1.min.js" type="text/javascript" charset="utf-8"></script>
	</head>
  </xsl:template>

  <xsl:template name="editBody">
	<body>
	<div id="editor"><xsl:value-of select="/edit/contents"/></div>
	<xform action="/edit" name="myform" method="POST">
	<!--input type="submit" value="Save" /-->
	<button onclick="saveFile()">Save</button>
    <script>
    var editor = ace.edit("editor");
    editor.setTheme("ace/theme/monokai");
    editor.getSession().setMode("ace/mode/golang");
	function saveFile(){alert("Saving File");
		$.post("/savingfile", "DATA GOES HERE")
	}
	
	</script>
	</xform>
</body>
</xsl:template>
</xsl:stylesheet>

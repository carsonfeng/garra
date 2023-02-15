export TAGNAME := 1.1.5

tag:
	git tag release-$(TAGNAME) -m $(TAGNAME)
	git push origin release-$(TAGNAME)
export TAGNAME := 1.1.6

tag:
	git tag release-$(TAGNAME) -m $(TAGNAME)
	git push origin release-$(TAGNAME)
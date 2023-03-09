export TAGNAME := 1.1.9.1

tag:
	git tag release-$(TAGNAME) -m $(TAGNAME)
	git push origin release-$(TAGNAME)
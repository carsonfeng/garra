export TAGNAME := 1.0.7

tag:
	git tag release-$(TAGNAME) -m $(TAGNAME)
	git push origin release-$(TAGNAME)
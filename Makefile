PREFIX        ?=  /usr
RM            ?=  rm -f
INSTALL_DIR   ?=  install -m755 -d
INSTALL_PROG  ?=  install -m755
INSTALL_FILE  ?=  install -m644

all:
	@echo Run \'make install\' to install prt-stuff.

install:
	$(INSTALL_DIR) $(DESTDIR)$(PREFIX)/bin
	$(INSTALL_DIR) $(DESTDIR)/etc
	$(INSTALL_DIR) $(DESTDIR)$(PREFIX)/share/fish/completions
	$(INSTALL_PROG) depls $(DESTDIR)$(PREFIX)/bin/depls
	$(INSTALL_PROG) depmk $(DESTDIR)$(PREFIX)/bin/depmk
	$(INSTALL_PROG) prtloc $(DESTDIR)$(PREFIX)/bin/prtloc
	$(INSTALL_PROG) prtprint $(DESTDIR)$(PREFIX)/bin/prtprint
	$(INSTALL_FILE) config/config $(DESTDIR)/etc/prt.conf
	$(INSTALL_FILE) completions/depls.fish $(DESTDIR)$(PREFIX)/share/fish/completions/depls.fish
	$(INSTALL_FILE) completions/depmk.fish $(DESTDIR)$(PREFIX)/share/fish/completions/depmk.fish
	$(INSTALL_FILE) completions/prtloc.fish $(DESTDIR)$(PREFIX)/share/fish/completions/prtloc.fish
	$(INSTALL_FILE) completions/prtprint.fish $(DESTDIR)$(PREFIX)/share/fish/completions/prtprint.fish

uninstall:
	$(RM) $(DESTDIR)$(PREFIX)/bin/depls
	$(RM) $(DESTDIR)$(PREFIX)/bin/depmk
	$(RM) $(DESTDIR)$(PREFIX)/bin/prtloc
	$(RM) $(DESTDIR)$(PREFIX)/bin/prtprint
	$(RM) $(DESTDIR)/etc/prt.conf
	$(RM) $(DESTDIR)$(PREFIX)/share/fish/completions/depls.fish
	$(RM) $(DESTDIR)$(PREFIX)/share/fish/completions/depmk.fish
	$(RM) $(DESTDIR)$(PREFIX)/share/fish/completions/prtloc.fish
	$(RM) $(DESTDIR)$(PREFIX)/share/fish/completions/prtprint.fish
	
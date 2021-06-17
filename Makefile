VER=0.9.6
PREFIX=/usr/share
DLURL="https://humungus.tedunangst.com/r/honk/d/honk-$(VER).tgz"
MANDIR="$(DESTDIR)$(PREFIX)/man/man"
DOCS="honk-$(VER)/docs"
SYSTEMD=1

getsrc:
	@if ! test -e honk-$(VER).tgz; then \
		$(MAKE) clean; \
		printf "Getting source for honk version $(VER)...\n"; \
		curl --progress-bar -fOJL $(DLURL); \
	fi
	@tar xzf honk-$(VER).tgz; \

rsstitle: rsstitle.patch
	$(MAKE) getsrc;
	cd honk-$(VER)/ && patch -Nbp1 < ../rsstitle.patch;

i18n: i18n.patch i18n.go
	$(MAKE) getsrc;
	cp i18n.go honk-$(VER)/;
	cp i18n_views/* honk-$(VER)/views/;
	cd honk-$(VER) && patch -Nbp1 < ../i18n.patch;
	
altnavbar: altnavbar.patch
	$(MAKE) getsrc;
	cp i18n_views/local.css honk-$(VER)/views;
	cd honk-$(VER) && patch -Nbp1 < ../altnavbar.patch;

i18naltnavbar: altnavbar_i18n.patch
	$(MAKE) getsrc;
	$(MAKE) i18n;
	cp i18n_views/local.css honk-$(VER)/views;
	cd honk-$(VER) && patch -Nbp1 < ../altnavbar_i18n.patch;

build: honk-$(VER)/
	cd honk-$(VER) && go build -mod=`ls -d vendor 2> /dev/null` -o honk;
	
install: honk.service honkinit honk-$(VER)/
	$(MAKE) build;
	install -Dm755 "honk-$(VER)/honk" -t "$(DESTDIR)/usr/bin/";
	install -Dm644 honk-$(VER)/views/* -t "$(DESTDIR)$(PREFIX)/honk/views/";
	install -Dm644 $(DOCS)/* -t "$(DESTDIR)$(PREFIX)/honk/docs/";
	gzip -k -f $(DOCS)/*.{1,3,5,7,8};
	install -Dm644 $(DOCS)/honk.1.gz -t $(MANDIR)1/;
	install -Dm644 $(DOCS)/honk.3.gz -t $(MANDIR)3/;
	install -Dm644 $(DOCS)/honk.5.gz -t $(MANDIR)5/;
	install -Dm644 $(DOCS)/honk.8.gz -t $(MANDIR)8/;
	install -Dm644 $(DOCS)/activitypub.7.gz $(MANDIR)7/honk_activitypub.7.gz;
	install -Dm644 $(DOCS)/hfcs.1.gz $(MANDIR)1/honk_hfcs.1.gz;
	install -Dm644 $(DOCS)/intro.1.gz $(MANDIR)1/honk_intro.1.gz;
	install -Dm644 $(DOCS)/vim.3.gz $(MANDIR)3/honk_vim.3.gz;	
	install -Dm644 "honk-$(VER)"/LICENSE -t "$(DESTDIR)$(PREFIX)/licenses/honk/";
	@if test -n $(SYSTEMD); #if systemd
	    install -Dm644 honk.service -t "$(DESTDIR)/usr/lib/systemd/system/";\
	    systemctl daemon-reload;\
	fi;
	mandb -q > /dev/null;
	if ! getent passwd honk >/dev/null; then \
        useradd -r -d $(DESTDIR)$(PREFIX)/honk honk; \
    fi
	chown honk:honk -R "$(DESTDIR)$(PREFIX)/honk";
	@if ! test -e $(PREFIX)/honk/honk.db; then \
       sh honkinit;\
    fi
	@printf "\nDone\n";

uninstall: $(DESTDIR)$(PREFIX)/honk;
	@if test -n $(SYSTEMD); #if systemd
	    systemctl stop honk;
	fi;
	userdel honk;
	rm -rf $(DESTDIR)$(PREFIX)/honk;
	rm -f $(DESTDIR)/usr/lib/systemd/system/honk.service;
	rm -f $(DESTDIR)/usr/bin/honk;
	find $(DESTDIR)$(PREFIX)/man/ -name "honk*.gz" -exec rm -f {} \;;

clean:
	rm -rf honk honk-$(VER)*;

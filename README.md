# honk_custom

[download here](https://git.les-miquelots.net/honk_custom/snapshot/honk_custom-master.zip)

Custom html templates and patches for the
[honk](https://humungus.tedunangst.com/r/honk) federated
ActivityPub server.
These are mainly related to the UX, but patching of the Go files
might happen if I feel the need.

Branches are created based on released version numbers.

# i18n

The patch `03_bloated_i18n_implementation.patch` works on
`web.go` in the latest changeset
of honk, [dca9f49c629f](https://humungus.tedunangst.com/r/honk/v/dca9f49c629f).

The translated strings are contained in the `i18n.go` and
`views/i18n/honkpage.js` files. Contact me if you need
help to add translations or notify me for problems, suggestions or
improvements.

Command line to do after adding translations to compile
and run the i18n'd honk should ressemble this:
```
cp patches/03_bloated_i18n_implementation.patch ../honk/
cp views/i18n/* ../honk/views/
cd ../honk
patch -b < 03_bloated_i18n_implementation.patch
make all
./honk
```

# Screenshots

Alternative navigation bar, less uses of the drop down menu:
[alt nav bar menu for honk](https://git.les-miquelots.net/honk_custom/plain/scrots/honk_altnavbar.png)

# honk_custom

[download here](https://git.les-miquelots.net/honk_custom/snapshot/honk_custom-master.zip)

Custom html templates and patches for the
[honk](https://humungus.tedunangst.com/r/honk) federated
ActivityPub server.
These are mainly related to the UX, but patching of the Go files
might happen if I feel the need.

# How to use

If you are already using Honk, backup
your `blob.db`, `honk.db` and `local.css`. For example, if you
want prepare the source with the rss patch:
```
make rsstitle
```
This will download the source tarball if not already done, untar it
and patch the source with the rsstitle.patch.
The patches are useable outside the Makefile. You just have to make
sure to put them in the root of the cloned honk repo or untarred source.

The `make rsstitle install` will :

- apply the rsstitle patch
- compile honk
- copy manuals to the relevant sections in `/usr/share/man`
- copy html template to `/usr/share/honk`
- created the **honk** user and group if not existing
- created a systemd service if `SYSTEMD=1` is defined in the Makefile
- initialize the honk database if not existing at `/usr/share/honk/honk.db`

# Patches

## i18n.patch : Translation of the honk user interface

- Adds an option to switch language, cookie-based (defaults to eng)
- Can also switch it in account settings

The patches `i18n.patch` works in the latest honk release, `0.9.6`, and
latest hg changeset of honk, [f74b9ce19463](https://humungus.tedunangst.com/r/honk/v/d/f74b9ce19463).

The translated strings are contained in the `i18n.go` and
`i18n_views/honkpage.js` files. Contact me if you need
help to add translations or notify me for problems, suggestions or
improvements.

## rsstitle.patch : More descriptive RSS title and description

Before the patch, the title and description of an user RSS Feed is like this:

- **title** : yourhandle honk
- **description**: yourhandle honk rss

After the patch:

- **title**: @yourhandle - honking from honk.club
- **description**: Honks from yourhandle@honk.club

I find it clearer.

## altnavbar.patch (and the i18n version): honking faster with less clicks

You need to shitpost fast ? Look no further, this will add a floating
navigation bar at the page's top with link to:
- /newhonk
- /front
- /home
- /xzone
- /@me
- /u/yourhandle (so your can see your own honks)
- /account

The `altnavbar_i18n.patch` is the same, but the labels are i18n'd. 

# Screenshots

A video of the i18n module and the navbar patch [here](https://partage.les-miquelots.net/img/honk_i18n.mp4)

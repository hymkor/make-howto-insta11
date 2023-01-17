**Please note:**

This is just a private tool to create a README.md from a template.

make-howto-insta11
==================

The make-howto-insta11 is the tool to output how to install our application with the scoop installer for README.md like this.

How to use
----------

```
cd YOUR-REPOSITORY
vim README.md
```

and type `:r!make-howto-insta11` to insert the section `Install` like this.

Install
-------

Download the binary package from [Releases](https://github.com/hymkor/make-howto-insta11/releases) and extract the executable.

### for scoop-installer

```
scoop install https://raw.githubusercontent.com/hymkor/make-howto-insta11/master/make-howto-insta11.json
```

or

```
scoop bucket add hymkor https://github.com/hymkor/scoop-bucket
scoop install make-howto-insta11
```

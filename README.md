# Cryptopals crypto challenges

| Link           | [the cryptopals crypto challenges](https://cryptopals.com/)                                                                                             |
| -------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Hilo de reddit | [What are some good projects in Go for an experienced dev?](https://www.reddit.com/r/golang/comments/17vdy90/what_are_some_good_projects_in_go_for_an/) |

# Set 1

## `Challenge 1` Convert hex to base64

The string:

```
49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d
```

Should produce:

```
SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t
```

So go ahead and make that happen. You'll need to use this code for the rest of the exercises.

> <b>Cryptopals Rule</b>  
> Always operate on raw bytes, never on encoded strings. Only use hex and base64 for pretty-printing.

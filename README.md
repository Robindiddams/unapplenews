# UnAppleNews

Tired of apple news openning when a co-worker posts an apple news article? Well now you can decode the redirect url and re-share it with a kind comment

```bash
# install from source:
$ go install github.com/Robindiddams/unapplenews/cmd/unapplenews

# run:
$ unapplenews https://apple.news/AArdKANhNR7KKXjwoNe9i-g
https://arstechnica.com/tech-policy/2020/10/googles-supreme-court-faceoff-with-oracle-was-a-disaster-for-google/

# OR open the url (mac only):
$ open `unapplenews https://apple.news/AArdKANhNR7KKXjwoNe9i-g`
```

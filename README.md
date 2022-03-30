# alt

An alternative to Markdown. (README ironically written in Markdown)

**alt** is a different approach to what Markdown is trying to achieve.

It's easiest to explain by showcasing it:

```alt
Hello ((world)), this (*is*) some (_formatted_) text! (fantastic)
```

Is the same as:

```md
Hello *world*, this **is** some __formatted__ text! (fantastic)
```

## alt text formatting reference

alt | html | markdown
----|------|---------
((...)) or (/.../) | `<i>...</i>` | \*...\* or \_...\_
(\*...\*) | `<b>...</b>` | \*\*...\*\* or \_\_...\_\_
(\_...\_) | `<u>...</u>` | *(Discord-only?)* \_\_...\_\_
(-...-) | `<s>...</s>` | \~\~...\~\~
(!...!) or (\|...\|) | `<mark>...</mark>` | *no equivalent*

### unimplemented formatting

alt | html | markdown
----|------|---------
(\`...\`) | `<code>...</code>` | \`...\`
('...') or (^...^) | `<super>...</super>` | *no equivalent*
(,...,) | `<sub>...</sub>` | *no equivalent*
("...") | `<blockquote>...</blockquote>` | *(Discord-only?)* > ...

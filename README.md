# ignore

## Name

*ignore* - ignores all queries (doesn't send any reply at all).

## Description

*ignore* ignores all queries, so CoreDNS can act e.g. as an authoritative DNS server without risking being a source of DNS amplification targets

## Syntax

~~~ txt
ignore
~~~

## Examples

Ignore all queries.

~~~ corefile
. {
    ignore
}
~~~

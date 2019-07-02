/*

Package legacysiva implements a go-borges library that uses siva files as its
storage backend. It has a 1-to-1 matching of locations and repositories and
no filtering or conversion is done to its references or objects.

It's meant to be used with siva files generated by borges. It's also a
read only implementation and does not support transactionality.

*/
package legacysiva

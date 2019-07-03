/*

This library abstracts read and write access to a set of go-git repositories.
It comes with several implementations to support different storage methods:

* plain: stored in the filesystem, supports transactions.

* siva: rooted repositories in siva files, supports transactions. These files can be generated with gitcollector.

* legacysiva: siva file generated by borges. This implementation only supports reading and does not support transactions.


This example lists the repositories downloaded by gitcollector:

	package main

	import (
		"fmt"
		"os"

		"github.com/src-d/go-borges"
		"github.com/src-d/go-borges/siva"
		"gopkg.in/src-d/go-billy.v4/osfs"
	)

	func main() {
		if len(os.Args) != 2 {
			fmt.Println("you need to provide the path of your siva files")
			os.Exit(1)
		}
		fs := osfs.New(os.Args[1])

		lib, err := siva.NewLibrary("library", fs, &siva.LibraryOptions{
			Bucket:        2,
			RootedRepo:    true,
			Transactional: true,
		})
		if err != nil {
			panic(err)
		}

		repos, err := lib.Repositories(borges.ReadOnlyMode)
		if err != nil {
			panic(err)
		}

		err = repos.ForEach(func(r borges.Repository) error {
			id := r.ID().String()
			head, err := r.R().Head()
			if err != nil {
				return err
			}

			fmt.Printf("repository: %v, HEAD: %v\n", id, head.Hash().String())
			return nil
		})
	}

More information:

* go-git: https://github.com/src-d/go-git

* rooted repositories: https://github.com/src-d/gitcollector#storing-repositories-using-rooted-repositories

* siva files: https://github.com/src-d/go-siva

* gitcollector: https://github.com/src-d/gitcollector

* borges: https://github.com/src-d/borges

*/
package borges

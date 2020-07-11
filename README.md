# goscp
Securely copy files from one server to another without the headaches maintaining bash script(s)

## Run
```
go install

goscp -i private_key_path -s source_file_path -d path_to_store_remote_file

```
You'll need to replace the paths with the correct file paths before using the command

Before copying a file, `goscp` checks whether the file already exists in the specified location(denoted by the -d flag) and if the file exists, we stop...else we continue to copy file.


## Yet to come!

1. Config.json
Specify different environments on the `config.json` file such that instead of passing all the fields as commandline arguments as above(which is cumbersome and error prone), you can just call the environment with your paths e.g

```
goscp prod_to_local
```

With this, `prod_to_local` is defined as(in the `config.json`):

```
{
	"prod_to_local": {
		"d": "/",
		"s": "/",
		"i": "/",
	}
}
```

2. Copy from one server to another via an intermediate environment. In this case, the intermediate environment could be having all the keys to access the two servers

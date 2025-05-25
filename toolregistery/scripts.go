package toolregistry

const installPostgresScript = `
cd {{ .TmpDir }}
curl -L https://ftp.postgresql.org/pub/source/v{{ .Version }}/postgresql-{{ .Version }}.tar.gz -o postgresql-{{ .Version }}.tar.gz
tar -zxvf postgresql-{{ .Version }}.tar.gz
cd postgresql-{{ .Version }}
./configure --prefix={{ .TmpDir }}/pgsql
make
make install
mv {{ .TmpDir }}/pgsql/bin/psql {{ .OutPath }}
`

const installMySQLScript = `
cd {{ .TmpDir }}
curl -L https://dev.mysql.com/get/Downloads/MySQL-{{ .Version }}/mysql-{{ .Version }}-linux-glibc2.12-x86_64.tar.gz -o mysql-{{ .Version }}.tar.gz
tar -zxvf mysql-{{ .Version }}.tar.gz
mv mysql-{{ .Version }}-linux-glibc2.12-x86_64/bin/mysql {{ .OutPath }}
`

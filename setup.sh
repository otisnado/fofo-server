#!/bin/bash
curl https://repo.maven.apache.org/maven2/org/springframework/boot/spring-boot-cli/$SPRING_CLI/spring-boot-cli-$SPRING_CLI-bin.zip -o springcli.zip
mkdir $HOME/.bin
unzip $HOME/springcli.zip
mv spring-$SPRING_CLI $HOME/.bin/
echo 'export SPRING_HOME=$HOME/.bin/spring-$SPRING_CLI' >> $HOME/.bashrc
echo 'export PATH=$PATH:$SPRING_HOME/bin' >> $HOME/.bashrc
source $HOME/.bashrc
rm $HOME/springcli.zip

for file in `find ./srv/ -type f -name "generate.sh"`;do
    sh $file
done
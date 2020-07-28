## Basis for cutemp.go when executing shell commands within  go lang


printf "["
for i in {1..10}
do
y=$(curl -s fact.user1.pks4u.com/fact/10 --write-out ",%{time_total}" | awk '{ for(i=1;i<=length($0);i++) if (substr($0,i,1)==",") { printf("%s",substr($0,i+1,5)); fflush(stdout);}}')
printf "[ "$i","$y" ]"
if [ $i -lt 10 ] 
then printf ","
fi
done
echo "]"


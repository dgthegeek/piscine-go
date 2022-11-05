num =  $(head -n 179 streets/Buckingham_place | tail -n i | cut -d '#' -f2)
echo $num
cat interviews/interview-$num
echo $MAIN_SUSPECT
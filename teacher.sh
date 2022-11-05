number =  $(head -n 179 streets/Buckingham_place | tail -n i | cut -d '#' -f2)
echo $number
cat interviews/interview-$number
echo $MAIN_SUSPECT
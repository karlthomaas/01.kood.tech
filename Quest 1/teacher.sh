inter_id=$(head -n 179 streets/Buckingham_Place | tail -n 1 | cut -f 2 -d '#' )
echo $inter_id
cat interviews/interview-$inter_id
echo $MAIN_SUSPECT

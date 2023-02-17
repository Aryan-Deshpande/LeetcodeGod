package graph


func graph(sr int, sc int, image [][]int, newcolor int) {
	if(image[sr][sc] == newcolor){return image;}
	fill(sr,sc,image,image[sr][sc],newcolor, image)
	return image;
}

func fill(sr int, sc int, color int, newcolor int, image [][]int){
	if(sc < 0 || sr < 0 || sr >= len(image) || sc >= len(image[0] || image[sr][sc]!=color)){return;}

	fill(sr+1,sc,color,newcolor,image)
	fill(sr,sc+1,color,newcolor,image)
	fill(sr-1,sc,color,newcolor,image)
	fill(sr,sc-1,color,newcolor,image)
}
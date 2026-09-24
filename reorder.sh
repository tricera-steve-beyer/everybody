goreorder reorder --diff ./ > reorder.patch
patch -p1  < ./reorder.patch
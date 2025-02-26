Key points:

Notice the pointer receiver *Matrix. 
This allows the method to modify the original struct's elements field, rather than a copy.
The method reads m.rows and m.columns to define the dimensions of the matrix dynamically.
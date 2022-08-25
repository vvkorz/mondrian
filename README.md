# Piet Mondrian like pictures in Go

## initial idea of an algorithm

- There are "color rectangles" and "line rectangles"
- Start with one color rectangle in some data structure (Array)
- [1] given function that takes a color rectangle and lambda cost and partitions it in 2 rectangles (use some distribution like poisson). The function returns 2 rectangels and the leftover costs
- pass initial color rectangle to function [1] that returns 2 color rectangles
- Replace initial color rectangle in the data structure (array) with 2 returned color rectangles
- [2] repeat recursively for 2 color rectangles.
- Stop if negative costs
- At this point we have a collection of color rectangles.
- [3] given function that takes collection of color rectangles and returns lines rectangles
-
- [4] given function that takes color rectangles, colors and mask and returns an image.Image assigning colors at random through masks 
- Call function [4] for color rectangles creating image 1
- Call function [4] for line rectangles creating image 2
- Overlay 2 images


## inspiration references

- [Mondrian Process](https://citeseerx.ist.psu.edu/viewdoc/download?doi=10.1.1.564.8410&rep=rep1&type=pdf)
- [GoMondrian](https://github.com/8lall0/GoMondrian)
- [generativeart](https://github.com/jdxyw/generativeart)

# Authors

- [pythonmonty](https://github.com/pythonmonty)
- [vvkorz](https://github.com/vvkorz)
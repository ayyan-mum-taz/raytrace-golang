
## What This Project Does

This is a  3D Ray Tracer implemented from scratch in Go. A ray tracer is a computer graphics technique that simulates how light travels through a scene, an advancement over rasterization which is how basically all real time graphics used to be rendered until NVIDIA's RTX cards. Ray tracing calculates the path of light rays to accurately model it. 

- **Reflection** (mirror-like surfaces)
- **Refraction** (glass/transparent materials)
- **Diffuse scattering** (matte surfaces)
- **Depth of field** (camera focus blur)
- **Anti-aliasing** (smooth edges)


## What Is Rendered

The raytracer renders a scene with:
- **Diffuse sphere**: (Lambertian)
- **Ground Plane**: Large sphere simulating a floor
- **Metallic sphere**: Metallic material with reflection
- **Glass sphere**: Dielectric material with refraction (hollow center creates a "bubble" effect)
- **Sky gradient**: Blue to white background

The output generates a PPM image file showing the rendered scene with realistic lighting, reflections, and refractions.


## How I accomplished this

I was only able to do this with the help of Peter Shirley's ray tracing tutorials which can be found at https://raytracing.github.io. 

No prebuilt graphics libraries were used, this was done with just math/algorithms. Implements real physics in reflection, refraction, and scattering.

## Explanation of directory

1. **Vector Mathematics** (`vector/`)
   - These contain the 3D vector operations (dot product, cross product, normalization) used for geometric calculations

2. **Ray Casting** (`ray/`)
   - Represents light rays with origin and direction, and calculates intersection points with objects

3. **Scene Objects** (`hitable/`)
   - Sphere geometry with ray-sphere intersection math

4. **Material System** (`material/`)
   - Lambertian: Diffuse/matte materials (absorbs and scatters light)
   - Metal: Reflective surfaces with configurable fuzziness
   - Dielectric: Transparent materials (glass) with refraction and Fresnel reflection
   - Implements physically-based rendering principles

5. **Camera System** (`camera/`)
   - Configurable field of view, aspect ratio

6. **Rendering Engine** (`main.go`)
   - Recursive ray tracing (up to 50 bounces)
   - Multi-sample anti-aliasing (100 samples per pixel)
   - Gamma correction for proper color display
   - Progress tracking with goroutines

## Some concepts used

Uses quadratic formula to find where rays hit spheres. Recursive ray tracing is important, primary rays and secondary rays for reflections and refractions, but there is a depth limit to prevent infinite recursion. Monte Carlo sampling (randomization algorithm) does multiple random samples per pixel, and averages results for smooth anti aliasing. Snel's law is used to calculate how light bends at material boundaries (refraction).



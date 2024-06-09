import sys
import os
from PIL import Image

def parse_args(args):
    image_dict = {}
    for arg in args:
        key, value = arg.split("=")
        image_dict[key] = value
    return image_dict

def process_image(image_paths, output_dir):
    print("python called")
    if not os.path.exists(output_dir):
        os.makedirs(output_dir)

    for image_path in image_paths:
        try:
            with Image.open(image_path) as img:
                # Process the image and save it
                processed_image_path = os.path.join(output_dir, "processed_" + os.path.basename(image_path))
                img.save(processed_image_path)
                print(f"Processed image saved as {processed_image_path}")
        except Exception as e:
            print(f"Failed to process {image_path}: {e}")

if __name__ == "__main__":
    if len(sys.argv) < 2:
        print("Usage: python3 process_image.py <key=value> <key=value> ...")
        sys.exit(1)

    output_dir = "./tmp"
    image_dict = parse_args(sys.argv[1:])
    image_paths = list(image_dict.values())
    process_image(image_paths, output_dir)

    # List the processed images
    print("\nList of processed images:")
    for filename in os.listdir(output_dir):
        print(filename)

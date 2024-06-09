import sys
import os
from PIL import Image

def process_image(image_paths, output_dir):
    print("python called")
    if not os.path.exists(output_dir):
        os.makedirs(output_dir)

    for image_path in image_paths:
        try:
            with Image.open(image_path) as img:
                # 処理後の画像を保存する
                processed_image_path = os.path.join(output_dir, "processed_" + os.path.basename(image_path))
                img.save(processed_image_path)
                print(f"Processed image saved as {processed_image_path}")
        except Exception as e:
            print(f"Failed to process {image_path}: {e}")

if __name__ == "__main__":
    if len(sys.argv) < 2:
        print("Usage: python3 process_image.py <image_path1> <image_path2> ...")
        sys.exit(1)

    output_dir = "./tmp"
    image_paths = sys.argv[1:]
    process_image(image_paths, output_dir)

    # 保存した画像の一覧を表示
    print("\nList of processed images:")
    for filename in os.listdir(output_dir):
        print(filename)

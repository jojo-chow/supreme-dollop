interface Props {
  previewUrl: string | null;
}

const ImageDisplay = ({ previewUrl }: Props) => {
  return (
    <div className="flex items-center justify-center w-48 h-48 bg-gray-100">
      {previewUrl ? (
        <img src={previewUrl} alt="Preview" className="w-full h-full object-cover" />
      ) : (
        <p>No Image Uploaded</p>
      )}
    </div>
  );
};

export default ImageDisplay;

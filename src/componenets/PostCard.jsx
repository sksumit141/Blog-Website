const PostCard = ({ post, onDelete }) => {
  return (
    <div className="bg-white p-4 rounded shadow">
      <h2 className="text-xl font-bold mb-2">{post.title}</h2>
      <p className="text-gray-600 mb-4">{post.content}</p>
      <div className="flex justify-between items-center">
        <span className="text-sm text-gray-500">By {post.author}</span>
        <button
          onClick={() => onDelete(post.id)}
          className="bg-red-500 text-white px-3 py-1 rounded hover:bg-red-600"
        >
          Delete
        </button>
      </div>
    </div>
  );
};

export default PostCard;
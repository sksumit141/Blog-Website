import { useState, useEffect } from "react";
import PostCard from "./componenets/PostCard";
import PostForm from "./componenets/PostForm";

const API_URL = "http://localhost:8080";

function App() {
  const [posts, setPosts] = useState([]);
  const [error, setError] = useState(null);

  const fetchPosts = async () => {
    try {
      const res = await fetch(`${API_URL}/posts`);
      if (!res.ok) throw new Error('Failed to fetch posts');
      const data = await res.json();
      setPosts(data.data);
    } catch (err) {
      setError(err.message);
      console.error('Error fetching posts:', err);
    }
  };

  useEffect(() => {
    fetchPosts();
  }, []);

  const handleDelete = async (id) => {
    try {
      const res = await fetch(`${API_URL}/posts/${id}`, {
        method: 'DELETE',
        headers: {
          'Content-Type': 'application/json'
        }
      });
      
      if (!res.ok) throw new Error('Failed to delete post');
      await fetchPosts();
    } catch (err) {
      setError(err.message);
      console.error('Error deleting post:', err);
    }
  };

  const handleCreate = async (post) => {
    try {
      const res = await fetch(`${API_URL}/posts`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(post),
      });

      if (!res.ok) throw new Error('Failed to create post');
      await fetchPosts();
    } catch (err) {
      setError(err.message);
      console.error('Error creating post:', err);
    }
  };

  return (
    <div className="p-6 max-w-4xl mx-auto">
      <h1 className="text-3xl font-bold mb-6 text-center">Go Blog</h1>
      {error && (
        <div className="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded mb-4">
          {error}
        </div>
      )}
      <PostForm onCreate={handleCreate} />
      <div className="grid grid-cols-1 md:grid-cols-2 gap-4 mt-6">
        {posts.map((post) => (
          <PostCard key={post.id} post={post} onDelete={handleDelete} />
        ))}
      </div>
    </div>
  );
}

export default App;
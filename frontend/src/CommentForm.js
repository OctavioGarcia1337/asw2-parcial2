import { useState } from "react";
import "./css/Item.css";


const CommentForm = ({
  handleSubmit,
  show,
  submitLabel,
  hasCancelButton = false,
  handleCancel,
  initialText = "",
}) => {
  const [text, setText] = useState(initialText);
  const isTextareaDisabled = text.length === 0;
  const onSubmit = (event) => {
    event.preventDefault();
    handleSubmit(text);
    setText("");
  };
  return (
    <form onSubmit={onSubmit}>
      <textarea className="comment-form-textarea" value={text} onChange={(e) => setText(e.target.value)}/>
      { show ? <button className="comment-form-button" disabled={isTextareaDisabled}> {submitLabel}</button> : void(0)}
      {hasCancelButton && (
      <button type="button" className="comment-form-button comment-form-cancel-button" onClick={handleCancel}>Cancel</button>
      )}
    </form>
  );
};

export default CommentForm;
